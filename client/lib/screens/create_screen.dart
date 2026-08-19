import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';

class CreateScreen extends StatefulWidget {
  final AppState state;

  const CreateScreen({super.key, required this.state});

  @override
  State<CreateScreen> createState() => _CreateScreenState();
}

class _CreateScreenState extends State<CreateScreen> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
    _controller.text = widget.state.player.flavorText;
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _submitText(String value) {
    setState(() {
      widget.state.player.flavorText = value;

      if (widget.state.player.inhabitKind.isForeigner) {
        widget.state.player.inhabitKind = InhabitKind.inhabitant;
        widget.state.inhabitants.add(widget.state.player);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('まどろみの水晶球'),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              SizedBox(height: 48),
              Text('水晶球が心に呼びかける。'),
              Text('夢の世界へいざないましょう。'),
              SizedBox(height: 24),
              Text('この世界で夢見るあなたは、どんな人ですか、どんな生きものですか。'),
              SizedBox(height: 24),
              Padding(
                padding: EdgeInsets.symmetric(horizontal: 48),
                child: TextField(
                  controller: _controller,
                  decoration: InputDecoration(hintText: '森の薬草取りの少女。'),
                  onSubmitted: (String value) => _submitText(value),
                  onChanged: (String value) => setState(() {}),
                  maxLines: null,
                ),
              ),
              SizedBox(height: 24),
              if (_controller.text == '')
                ElevatedButton(onPressed: null, child: Text('答える'))
              else
                ElevatedButton(
                  onPressed: () => _submitText(_controller.text),
                  child: Text('答える'),
                ),
              if (widget.state.player.flavorText != '') ...[
                SizedBox(height: 48),
                Text(widget.state.player.flavorText),
              ],
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('目をそらす'),
              ),
              SizedBox(height: 48),
            ],
          ),
        ),
      ),
    );
  }
}
