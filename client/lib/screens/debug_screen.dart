import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';

class DebugScreen extends StatefulWidget {
  final AppState state;

  const DebugScreen({super.key, required this.state});

  @override
  State<DebugScreen> createState() => _DebugScreenState();
}

class _DebugScreenState extends State<DebugScreen> {
  void _clearMyself() {
    setState(() => widget.state.player = Player.unnamed());
  }

  void _clearInhabitants() {
    setState(() {
      widget.state.player = Player.unnamed();
      widget.state.inhabitants = [];
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('汚れた道具箱'),
      ),

      body: Center(
        child: Column(
          mainAxisAlignment: .center,
          children: [
            ElevatedButton(
              onPressed: widget.state.player.inhabitKind.isForeigner
                  ? null
                  : () => _clearMyself(),
              child: const Text('自分を手放す'),
            ),

            const SizedBox(height: 48),

            ElevatedButton(
              onPressed: widget.state.inhabitants.isEmpty
                  ? null
                  : () => _clearInhabitants(),
              child: const Text('住人たちを追い出す'),
            ),

            const SizedBox(height: 48),

            ElevatedButton(
              onPressed: () => Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => const HelpScreen(page: HelpPage.debug),
                ),
              ),
              child: const Text('親切な鏡の精'),
            ),

            const SizedBox(height: 96),

            ElevatedButton(
              onPressed: () => Navigator.pushNamed(context, '/'),
              child: const Text('ふたを閉じる'),
            ),
          ],
        ),
      ),
    );
  }
}
