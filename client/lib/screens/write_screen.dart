import 'package:flutter/material.dart';

import 'package:client/api/commit_action.dart';
import 'package:client/api/ensure_session.dart';
import 'package:client/api/new_action.dart';
import 'package:client/api/image_action.dart';
import 'package:client/models/action.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/utils/image_url.dart';
import 'package:client/widgets/translucent_panel.dart';

class WriteScreen extends StatefulWidget {
  final AppState state;

  const WriteScreen({super.key, required this.state});

  @override
  State<WriteScreen> createState() => _WriteScreenState();
}

class _WriteScreenState extends State<WriteScreen> {
  PlayerAction _action = PlayerAction();

  late TextEditingController _inputController;
  String? _actionErrorMessage;
  late TextEditingController _outputController;
  bool _actionLoading = false;
  bool _imageLoading = false;
  String? _imageErrorMessage;

  @override
  void initState() {
    super.initState();
    _inputController = TextEditingController();
    _outputController = TextEditingController();
  }

  @override
  void dispose() {
    _inputController.dispose();
    _outputController.dispose();
    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _sync();
  }

  Future<void> _sync() async {
    if (await widget.state.sync()) {
      setState(() {});
    }

    if (widget.state.action != null) {
      _action = PlayerAction();
      _action.description = widget.state.action!.description;
      _action.imageId = widget.state.action!.imageId;
      _outputController.text = _action.description;

      setState(() {});
    }
  }

  Future<void> _loadImage() async {
    if (widget.state.restAiCalls < 1) {
      return;
    }
    setState(() {
      _imageLoading = true;
      _imageErrorMessage = null;
    });
    try {
      widget.state.incrementAiCalls();
      setState(() {});
      final imageId = await apiImageAction();

      setState(() {
        _action.imageId = imageId;
      });
    } catch (e) {
      setState(() => _imageErrorMessage = e.toString());
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _loadAction() async {
    if (widget.state.restAiCalls < 1) {
      return;
    }
    setState(() {
      _actionLoading = true;
      _actionErrorMessage = null;
    });
    try {
      await apiEnsureSession();

      widget.state.incrementAiCalls();
      setState(() {});
      final action = await apiNewAction(_inputController.text);
      if (action.error == '') {
        action.input = _inputController.text;

        setState(() {
          _outputController.text = action.description;
          _action = action;
        });

        _loadImage();
      } else {
        setState(() {
          _outputController.text = 'エラー:\n${action.error}';
        });
      }
    } catch (e) {
      setState(() => _actionErrorMessage = e.toString());
    } finally {
      setState(() => _actionLoading = false);
    }
  }

  Future<void> _commit() async {
    await apiCommitAction();

    widget.state.clear();
    await _sync();
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
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  if (widget.state.inhabitant) ...[
                    if (widget.state.day != null)
                      TranslucentPanel(
                        child: Column(
                          children: [
                            Text(formatDate(widget.state.day!)),
                            Text(
                              '住み着いて ${widget.state.day! - widget.state.player!.day + 1} 日目',
                            ),
                          ],
                        ),
                      ),

                    const SizedBox(height: 24),

                    TranslucentPanel(child: const Text('日記に書いてみよう。')),

                    const SizedBox(height: 24),

                    TranslucentPanel(child: const Text('今日は何をして過ごしましたか？')),
                  ],

                  const SizedBox(height: 24),

                  if (widget.state.inhabitant)
                    Padding(
                      padding: EdgeInsets.all(
                        MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
                      ),
                      child: TranslucentPanel(
                        child: TextField(
                          controller: _inputController,
                          decoration: const InputDecoration(
                            hintText: '(今日、何をしたか書いてみよう。)',
                          ),
                          onChanged: (String value) => setState(() {}),
                          maxLines: null,
                          maxLength: 140,
                          readOnly:
                              _actionLoading ||
                              _imageLoading ||
                              widget.state.committed,
                        ),
                      ),
                    ),

                  const SizedBox(height: 24),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed:
                          (_inputController.text.isEmpty ||
                              _actionLoading ||
                              _imageLoading ||
                              widget.state.committed)
                          ? null
                          : _loadAction,
                      child: const Text('日記に書く'),
                    ),
                  ),

                  if (_actionErrorMessage != null) const SizedBox(height: 12),
                  if (_actionErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _actionErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

                  if (_actionLoading)
                    const CircularProgressIndicator()
                  else if (_action.hasDescription) ...[
                    const SizedBox(height: 48),

                    TranslucentPanel(
                      child: const Text('あなたの言葉は夢に映され、こうなりました。'),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: TextField(
                        controller: _outputController,
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                    const SizedBox(height: 24),

                    if (_imageLoading)
                      const Column(
                        children: [
                          TranslucentPanel(child: Text('情景を映し出しています…')),
                          CircularProgressIndicator(),
                        ],
                      )
                    else if (_action.imageId != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.network(imageUrl(_action.imageId!)),
                          ),
                        ),
                      ),

                    if (_imageErrorMessage != null) const SizedBox(height: 12),
                    if (_imageErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _imageErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed:
                          (_actionLoading ||
                              _imageLoading ||
                              widget.state.committed)
                          ? null
                          : () async {
                              await _commit();

                              if (!context.mounted) return;

                              Navigator.pushNamed(context, '/home');
                            },
                      child: const Text('これでよし'),
                    ),
                  ],

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: (_actionLoading || _imageLoading)
                        ? null
                        : () => Navigator.pushNamed(context, '/home'),
                    child: const Text('ペンを置く'),
                  ),

                  const SizedBox(height: 96),
                ],
              ),
            ),
          ),

          Positioned(
            bottom: 24,
            left: 8,
            right: 8,
            child: Row(
              children: [
                for (int i = 0; i < widget.state.restAiCalls; i++)
                  Image(image: AssetImage('assets/images/fruit.webp')),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
