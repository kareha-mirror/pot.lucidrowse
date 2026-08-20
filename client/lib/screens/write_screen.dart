import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
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

  void _submitText(String value) {
    setState(() {
      widget.state.player.action.raw = value;
      widget.state.player.action.filtered = value;
      widget.state.player.actions.add(widget.state.player.action);
      _controller.text = '';
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('やわらかい羽ペン'),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              SizedBox(height: 48),
              if (!widget.state.player.inhabit.isForeigner) ...[
                Text('日記に書いてみよう。'),
                SizedBox(height: 24),
                Text('今日は何をして過ごしましたか？'),
              ],
              SizedBox(height: 24),
              if (!widget.state.player.inhabit.isForeigner)
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 48),
                  child: TextField(
                    controller: _controller,
                    decoration: InputDecoration(hintText: '冒険者の泉で魚釣りをしてみた。'),
                    onSubmitted: (String value) => _submitText(value),
                    onChanged: (String value) => setState(() {}),
                    maxLines: null,
                    maxLength: 140,
                  ),
                ),
              SizedBox(height: 24),
              if (_controller.text == '')
                ElevatedButton(onPressed: null, child: Text('日記に書く'))
              else
                ElevatedButton(
                  onPressed: () => _submitText(_controller.text),
                  child: Text('日記に書く'),
                ),
              if (widget.state.player.action.filtered != '') ...[
                SizedBox(height: 48),
                Text(widget.state.player.action.filtered),
              ],
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('ペンを置く'),
              ),
              SizedBox(height: 48),
            ],
          ),
        ),
      ),
    );
  }
}
