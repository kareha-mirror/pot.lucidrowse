import 'package:flutter/material.dart';

import 'package:client/api/day.dart';
import 'package:client/api/hello.dart';
import 'package:client/api/next_day.dart';
import 'package:client/models/player.dart';
import 'package:client/screens/help_screen.dart';
import 'package:client/state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/widgets/translucent_panel.dart';

class DebugScreen extends StatefulWidget {
  final AppState state;

  const DebugScreen({super.key, required this.state});

  @override
  State<DebugScreen> createState() => _DebugScreenState();
}

class _DebugScreenState extends State<DebugScreen> {
  String? _message;

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

  void _nextDay() async {
    setState(() {
      widget.state.player.committed = false;
      widget.state.player.action = PlayerAction();
    });

    try {
      final result = await apiNextDay();
      setState(() => widget.state.day = result['day']);
    } catch (e) {
      setState(() => _message = e.toString());
    }
  }

  void _clearMyself() {
    setState(() => widget.state.player = Player());
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
                    onPressed: _nextDay,
                    child: const Text('明日まで寝る'),
                  ),

                  const SizedBox(height: 48),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed: !widget.state.player.inhabitant
                          ? null
                          : () => _clearMyself(),
                      child: const Text('自分を手放す'),
                    ),
                  ),

                  const SizedBox(height: 48),

                  const SizedBox(height: 48),

                  ElevatedButton(
                    onPressed: () async {
                      try {
                        final result = await apiHello();
                        setState(() => _message = result['message']);
                      } catch (e) {
                        setState(() => _message = e.toString());
                      }
                    },
                    child: const Text('夢の世界に呼びかける'),
                  ),

                  const SizedBox(height: 12),

                  if (_message != null)
                    TranslucentPanel(child: Text(_message!)),

                  const SizedBox(height: 48),

                  helpButton(context, HelpPage.debug),

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: const Text('ふたを閉じる'),
                  ),

                  const SizedBox(height: 512),

                  SizedBox(
                    width: 300,
                    child: SwitchListTile(
                      title: TranslucentPanel(
                        child: Center(child: const Text('時空の裂け目')),
                      ),
                      value: widget.state.override,
                      onChanged: (value) =>
                          setState(() => widget.state.override = value),
                    ),
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
