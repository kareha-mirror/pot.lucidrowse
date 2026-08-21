import 'package:flutter/material.dart';

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
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
    _controller.text = widget.state.player.action.raw;
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _submitRaw(String value) {
    setState(() {
      final action = widget.state.player.action;

      action.raw = value;
      action.filtered = value;

      action.day = widget.state.day;
    });
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
                          controller: _controller,
                          decoration: const InputDecoration(
                            hintText: '(今日、何をしたか書いてみよう。)',
                          ),
                          onSubmitted: (String value) => _submitRaw(value),
                          onChanged: (String value) => setState(() {}),
                          maxLines: null,
                          maxLength: 140,
                        ),
                      ),
                    ),

                  const SizedBox(height: 24),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: _controller.text == ''
                          ? null
                          : () => _submitRaw(_controller.text),
                      child: const Text('日記に書く'),
                    ),
                  ),

                  if (widget.state.player.action.hasFiltered) ...[
                    const SizedBox(height: 48),

                    TranslucentPanel(
                      child: const Text('あなたの言葉は夢に映され、こうなりました。'),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: Text(widget.state.player.action.filtered),
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
