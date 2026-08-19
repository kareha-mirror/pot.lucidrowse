import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';

class DebugScreen extends StatefulWidget {
  final AppState state;

  const DebugScreen({super.key, required this.state});

  @override
  State<DebugScreen> createState() => _DebugScreenState();
}

class _DebugScreenState extends State<DebugScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('汚れた道具箱'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            SizedBox(height: 48),
            if (widget.state.player.inhabitKind.isForeigner)
              ElevatedButton(onPressed: null, child: Text('自分を手放す'))
            else
              ElevatedButton(
                onPressed: () =>
                    setState(() => widget.state.player = Player.unnamed()),
                child: Text('自分を手放す'),
              ),
            SizedBox(height: 48),
            if (widget.state.inhabitants.isEmpty)
              ElevatedButton(onPressed: null, child: Text('住人たちを追い出す'))
            else
              ElevatedButton(
                onPressed: () => setState(() {
                  widget.state.player = Player.unnamed();
                  widget.state.inhabitants = [];
                }),
                child: Text('住人たちを追い出す'),
              ),
            SizedBox(height: 96),
            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: Text('ふたを閉じる'),
            ),
            Spacer(),
          ],
        ),
      ),
    );
  }
}
