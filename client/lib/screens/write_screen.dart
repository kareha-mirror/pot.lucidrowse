import 'package:flutter/material.dart';

import 'package:client/api/action.dart';
import 'package:client/api/action_image.dart';
import 'package:client/api/image.dart';
import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/widgets/translucent_panel.dart';

class WriteScreen extends StatefulWidget {
  final AppState state;

  const WriteScreen({super.key, required this.state});

  @override
  State<WriteScreen> createState() => _WriteScreenState();
}

class _WriteScreenState extends State<WriteScreen> {
  late TextEditingController _inputController;
  late TextEditingController _outputController;
  bool _textLoading = false;
  bool _imageLoading = false;

  @override
  void initState() {
    super.initState();
    _inputController = TextEditingController();
    _inputController.text = widget.state.player.action.raw;
    _outputController = TextEditingController();
  }

  @override
  void dispose() {
    _inputController.dispose();
    _outputController.dispose();
    super.dispose();
  }

  Future<void> _fetchImage(String id) async {
    setState(() => _imageLoading = true);
    try {
      final image = await apiActionImage(id);

      setState(() {
        widget.state.player.action.image = image;
      });
    } catch (e) {
      // TODO
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _fetchAction() async {
    setState(() => _textLoading = true);
    try {
      final filtered = await apiAction(
        widget.state.player.id,
        _inputController.text,
      );
      final a = PlayerAction();

      a.raw = _inputController.text;
      a.day = widget.state.day;

      a.filtered = filtered['text'];

      setState(() {
        widget.state.player.action = a;
        _outputController.text = widget.state.player.action.filtered;
      });

      _fetchImage(filtered['id']);
    } catch (e) {
      // TODO
    } finally {
      setState(() => _textLoading = false);
    }
  }

  void _commit() {
    setState(() {
      final player = widget.state.player;
      player.committed = true;
    });
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
              image: AssetImage(
                widget.state.player.committed
                    ? 'assets/images/night.webp'
                    : 'assets/images/home.webp',
              ),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  if (!widget.state.player.isForeigner) ...[
                    TranslucentPanel(
                      child: Column(
                        children: [
                          Text(formatDate(widget.state.day)),
                          Text(
                            'この世界に住んで ${widget.state.day - widget.state.player.settled + 1} 日目。',
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

                  if (!widget.state.player.isForeigner)
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
                          readOnly: _textLoading || _imageLoading,
                        ),
                      ),
                    ),

                  const SizedBox(height: 24),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed:
                          (_inputController.text.isEmpty ||
                              _textLoading ||
                              _imageLoading)
                          ? null
                          : _fetchAction,
                      child: const Text('日記に書く'),
                    ),
                  ),

                  if (_textLoading)
                    const CircularProgressIndicator()
                  else if (widget.state.player.action.hasFiltered) ...[
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
                          TranslucentPanel(child: const Text('情景を映し出しています…')),
                          const CircularProgressIndicator(),
                        ],
                      )
                    else if (widget.state.player.action.image != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.memory(
                              widget.state.player.action.image!,
                            ),
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed: () {
                        _commit();
                        Navigator.pushNamed(context, '/home');
                      },
                      child: const Text('これでよし'),
                    ),
                  ],

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pop(context),
                    child: const Text('ペンを置く'),
                  ),

                  const SizedBox(height: 96),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
