import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/widgets/translucent_panel.dart';

class DebugScreen extends StatefulWidget {
  final AppState state;

  const DebugScreen({super.key, required this.state});

  @override
  State<DebugScreen> createState() => _DebugScreenState();
}

class _DebugScreenState extends State<DebugScreen> {
  void _nextDay() {
    setState(() {
      for (final inhabitant in widget.state.inhabitants) {
        if (inhabitant.committed) {
          inhabitant.actions.add(inhabitant.action);
        }
        inhabitant.action = PlayerAction();
      }

      widget.state.day++;
    });
  }

  void _clearMyself() {
    setState(() => widget.state.player = Player());
  }

  void _clearInhabitants() {
    setState(() {
      widget.state.player = Player();
      widget.state.inhabitants = [];
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
            child: const Image(
              image: AssetImage('assets/images/debug.webp'),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  TranslucentPanel(
                    child: Column(
                      children: [
                        Text(formatDate(widget.state.day)),
                        Text('夢路開通 ${widget.state.day + 1} 日目'),
                      ],
                    ),
                  ),
                  const SizedBox(height: 12),
                  ElevatedButton(
                    onPressed: () => _nextDay(),
                    child: const Text('明日まで寝る'),
                  ),

                  const SizedBox(height: 48),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: widget.state.player.isForeigner
                          ? null
                          : () => _clearMyself(),
                      child: const Text('自分を手放す'),
                    ),
                  ),

                  const SizedBox(height: 48),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: widget.state.inhabitants.isEmpty
                          ? null
                          : () => _clearInhabitants(),
                      child: const Text('住人たちを追い出す'),
                    ),
                  ),

                  const SizedBox(height: 48),

                  helpButton(context, HelpPage.debug),

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: const Text('ふたを閉じる'),
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
