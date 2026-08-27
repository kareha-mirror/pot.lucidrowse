import 'package:flutter/material.dart';

import 'package:client/api/commit_action.dart';
import 'package:client/api/day.dart';
import 'package:client/api/new_action.dart';
import 'package:client/api/image_action.dart';
import 'package:client/const.dart';
import 'package:client/models/player.dart';
import 'package:client/state.dart';
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
  bool _actionLoading = false;
  bool _imageLoading = false;

  @override
  void initState() {
    super.initState();
    _inputController = TextEditingController();
    _inputController.text = widget.state.player.action.input;
    _outputController = TextEditingController();
    _outputController.text = widget.state.player.action.description;
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

    _loadBackground();
  }

  Future<void> _loadBackground() async {
    final result = await apiDay();

    if (!mounted) return;

    setState(() => widget.state.day = result['day']);
  }

  Future<void> _loadImage() async {
    setState(() => _imageLoading = true);
    try {
      final result = await apiImageAction(widget.state.player.id);

      setState(() {
        widget.state.player.action.imageId = result['image-id'];
      });
    } catch (e) {
      // TODO
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _loadAction() async {
    setState(() => _actionLoading = true);
    try {
      final result = await apiNewAction(
        widget.state.player.id,
        _inputController.text,
      );
      if (result['error'] == '') {
        final action = PlayerAction();

        action.input = _inputController.text;

        action.description = result['description'];

        action.day = widget.state.day;

        setState(() {
          _outputController.text = action.description;
          widget.state.player.action = action;
        });

        _loadImage();
      } else {
        setState(() {
          _outputController.text = 'エラー:\n${result['error']}';
        });
      }
    } catch (e) {
      // TODO
    } finally {
      setState(() => _actionLoading = false);
    }
  }

  Future<void> _commit() async {
    await apiCommitAction(widget.state.player.id);

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
                  if (widget.state.player.inhabitant) ...[
                    TranslucentPanel(
                      child: Column(
                        children: [
                          Text(formatDate(widget.state.day)),
                          Text(
                            'この世界に住んで ${widget.state.day - widget.state.player.day + 1} 日目。',
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

                  if (widget.state.player.inhabitant)
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
                          readOnly: _actionLoading || _imageLoading,
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
                              widget.state.player.committed)
                          ? null
                          : _loadAction,
                      child: const Text('日記に書く'),
                    ),
                  ),

                  if (_actionLoading)
                    const CircularProgressIndicator()
                  else if (widget.state.player.action.hasDescription) ...[
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
                    else if (widget.state.player.action.imageId != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.network(
                              '$apiBase/image/${widget.state.player.action.imageId!}',
                            ),
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed:
                          (_actionLoading ||
                              _imageLoading ||
                              widget.state.player.committed)
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
                        : () => Navigator.pop(context),
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
