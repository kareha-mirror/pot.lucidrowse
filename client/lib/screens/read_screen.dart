import 'dart:async';
import 'package:flutter/material.dart';

import 'package:client/api/api.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/image_url.dart';
import 'package:client/widgets/translucent_panel.dart';

class ReadScreen extends StatefulWidget {
  final AppState state;

  final String? playerId;

  const ReadScreen({super.key, required this.state, this.playerId});

  @override
  State<ReadScreen> createState() => _ReadScreenState();
}

class _ReadScreenState extends State<ReadScreen> {
  Timer? _syncTimer;

  int _page = 1;
  bool _hasNext = false;

  bool _initialized = false;
  String? _actionsErrorMessage;
  List<dynamic> _actions = [];

  late ScrollController _scrollController;

  @override
  void initState() {
    super.initState();

    _syncTimer = .periodic(const .new(hours: 1), (_) => _sync());

    _scrollController = .new();
  }

  @override
  void dispose() {
    _scrollController.dispose();

    _syncTimer?.cancel();

    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _loadActions();
  }

  Future<void> _sync() async {
    if (await widget.state.sync()) {
      if (!mounted) return;

      setState(() {});
    }
  }

  Future<void> _loadActions() async {
    await _sync();
    setState(() => _actions = []);
    try {
      final Map<String, dynamic> result;
      if (widget.playerId == null) {
        if (widget.state.player == null) {
          return;
        }
        result = await api.listActions(widget.state.player!.id, _page);
      } else {
        result = await api.listActions(widget.playerId!, _page);
      }

      if (!mounted) return;

      setState(() {
        _actions = result['actions'] ?? [];
        _hasNext = result['has-next'] ?? false;
      });
    } catch (e) {
      setState(() => _actionsErrorMessage = e.toString());
    } finally {
      setState(() => _initialized = true);
    }
  }

  List<Widget> _actionList() {
    if (_initialized && _actions.isEmpty) {
      return [
        const SizedBox(height: 24),
        TranslucentPanel(child: const Text('まだ何も書かれてない。')),
      ];
    }

    List<Widget> list = [];
    for (final action in _actions) {
      list.add(const SizedBox(height: 24));
      list.add(
        TranslucentPanel(
          child: Column(
            children: [
              Text(action['date']),
              SizedBox(height: 8),
              Text(action['description']),
            ],
          ),
        ),
      );
      list.add(SizedBox(height: 8));
      if (action['image-id'] != null) {
        list.add(
          ConstrainedBox(
            constraints: const .new(maxWidth: 300),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: AspectRatio(
                aspectRatio: 1.0,
                child: Image.network(
                  imageUrl(action['image-id']),
                  frameBuilder:
                      (context, child, frame, wasSynchronouslyLoaded) {
                        if (wasSynchronouslyLoaded || frame != null) {
                          return child;
                        }

                        return const Center(child: CircularProgressIndicator());
                      },
                ),
              ),
            ),
          ),
        );
      }
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return PopScope(
      canPop: false,
      onPopInvokedWithResult: (didPop, result) {
        if (!didPop) {
          Navigator.pop(context);
        }
      },
      child: Scaffold(
        body: Stack(
          children: [
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(
                  widget.state.night
                      ? 'assets/images/night.webp'
                      : 'assets/images/home.webp',
                ),
                fit: .cover,
              ),
            ),

            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
              controller: _scrollController,
              child: Center(
                child: Column(
                  children: [
                    if (_actionsErrorMessage != null)
                      const SizedBox(height: 12),
                    if (_actionsErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _actionsErrorMessage!,
                          style: .new(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    if (_page > 1) SizedBox(height: 48),

                    if (_page > 1)
                      ElevatedButton(
                        onPressed: () async {
                          setState(() => _page--);
                          await _loadActions();
                          WidgetsBinding.instance.addPostFrameCallback((_) {
                            if (_scrollController.hasClients) {
                              _scrollController.jumpTo(
                                _scrollController.position.maxScrollExtent,
                              );
                            }
                          });
                        },
                        child: const Text('もっと後を読む'),
                      ),

                    if (_page > 1) SizedBox(height: 48),

                    ..._actionList(),

                    if (_hasNext) SizedBox(height: 48),

                    if (_hasNext)
                      ElevatedButton(
                        onPressed: () async {
                          setState(() => _page++);
                          await _loadActions();
                          _scrollController.jumpTo(0);
                          WidgetsBinding.instance.addPostFrameCallback((_) {
                            if (_scrollController.hasClients) {
                              _scrollController.jumpTo(0);
                            }
                          });
                        },
                        child: const Text('もっと前を読む'),
                      ),

                    SizedBox(height: 96),

                    ElevatedButton(
                      onPressed: () => Navigator.pop(context),
                      child: const Text('閉じる'),
                    ),

                    SizedBox(height: 96),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

Widget myDiaryButton(BuildContext context, AppState state) {
  return ElevatedButton(
    onPressed: () => Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => ReadScreen(state: state)),
    ),
    child: const Text('おしゃれな日記帳'),
  );
}

Widget diaryButton(BuildContext context, AppState state, String playerId) {
  return ElevatedButton(
    onPressed: () => Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ReadScreen(state: state, playerId: playerId),
      ),
    ),
    child: const Text('日記'),
  );
}
