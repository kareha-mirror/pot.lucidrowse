import 'package:flutter/material.dart';

import 'package:client/state/app_state.dart';

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
      player.action.committed = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('やわらかい羽ペン'),
      ),

      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Center(
          child: Column(
            children: [
              if (!widget.state.player.inhabit.isForeigner) ...[
                const Text('日記に書いてみよう。'),
                const SizedBox(height: 24),
                const Text('今日は何をして過ごしましたか？'),
              ],

              const SizedBox(height: 24),

              if (!widget.state.player.inhabit.isForeigner)
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 48),
                  child: TextField(
                    controller: _controller,
                    decoration: const InputDecoration(
                      hintText: '冒険者の泉で魚釣りをしてみた。',
                    ),
                    onSubmitted: (String value) => _submitRaw(value),
                    onChanged: (String value) => setState(() {}),
                    maxLines: null,
                    maxLength: 140,
                  ),
                ),

              const SizedBox(height: 24),

              ElevatedButton(
                onPressed: _controller.text == ''
                    ? null
                    : () => _submitRaw(_controller.text),
                child: const Text('日記に書く'),
              ),

              if (widget.state.player.action.filtered != '') ...[
                const SizedBox(height: 48),

                const Text('あなたの言葉は夢に映されこうなりました。'),

                const SizedBox(height: 24),

                Text(widget.state.player.action.filtered),

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
            ],
          ),
        ),
      ),
    );
  }
}
