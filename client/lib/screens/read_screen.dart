import 'package:flutter/material.dart';

import 'package:client/api/list_actions.dart';
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
  bool _initialized = false;
  String? _actionsErrorMessage;
  List<dynamic> _actions = [];

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _loadActions();
  }

  Future<void> _sync() async {
    if (await widget.state.sync()) {
      setState(() {});
    }
  }

  Future<void> _loadActions() async {
    await _sync();
    setState(() => _initialized = false);
    try {
      final Map<String, dynamic> result;
      if (widget.playerId == null) {
        if (widget.state.player == null) {
          return;
        }
        result = await apiListActions(widget.state.player!.id);
      } else {
        result = await apiListActions(widget.playerId!);
      }

      if (!mounted) return;

      setState(() => _actions = result['actions'] ?? []);
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
            constraints: const BoxConstraints(maxWidth: 300),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Image.network(imageUrl(action['image-id'])),
            ),
          ),
        );
      }
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          SizedBox(
            width: double.infinity,
            height: double.infinity,
            child: Image(
              /* TODO night
              image: AssetImage(
                _player.committed
                    ? 'assets/images/night.webp'
                    : 'assets/images/home.webp',
              ),
              */
              image: AssetImage('assets/images/home.webp'),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
            child: Center(
              child: Column(
                children: [
                  if (_actionsErrorMessage != null) const SizedBox(height: 12),
                  if (_actionsErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _actionsErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

                  ..._actionList(),

                  SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pop(context),
                    child: Text('閉じる'),
                  ),

                  SizedBox(height: 96),
                ],
              ),
            ),
          ),
        ],
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
