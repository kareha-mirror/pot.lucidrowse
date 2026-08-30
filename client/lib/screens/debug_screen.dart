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
  bool _initialized = false;
  String? _dayErrorMessage;
  bool _sleeping = false;
  String? _nextDayErrorMessage;
  String? _message;
  String? _helloErrorMessage;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _loadDay();
  }

  Future<void> _loadDay() async {
    try {
      final result = await apiDay();

      if (!mounted) return;

      setState(() {
        widget.state.day = result['day'];

        _initialized = true;
      });
    } catch (e) {
      setState(() => _dayErrorMessage = e.toString());
    }
  }

  void _nextDay() async {
    setState(() {
      _sleeping = true;
      _nextDayErrorMessage = null;
    });

    try {
      final result = await apiNextDay();

      if (result['error'] != '') {
        setState(() => _nextDayErrorMessage = result['error']);
      } else {
        setState(() {
          widget.state.day = result['day'];

          widget.state.player.committed = false;
          widget.state.player.action = PlayerAction();
        });
      }
    } catch (e) {
      setState(() => _nextDayErrorMessage = e.toString());
    } finally {
      setState(() => _sleeping = false);
    }
  }

  void _clearMyself() {
    setState(() => widget.state.player = Player());
  }

  void _hello() async {
    try {
      setState(() => _message = null);

      final result = await apiHello();

      setState(() => _message = result['message']);
    } catch (e) {
      setState(() => _helloErrorMessage = e.toString());
    }
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
                  if (_initialized)
                    TranslucentPanel(
                      child: Column(
                        children: [
                          Text(formatDate(widget.state.day)),
                          Text('夢路開通 ${widget.state.day + 1} 日目'),
                        ],
                      ),
                    ),

                  if (_dayErrorMessage != null) const SizedBox(height: 12),
                  if (_dayErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _dayErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

                  const SizedBox(height: 12),
                  ElevatedButton(
                    onPressed: _sleeping ? null : _nextDay,
                    child: const Text('明日まで寝る'),
                  ),

                  if (_sleeping) const CircularProgressIndicator(),

                  if (_nextDayErrorMessage != null) const SizedBox(height: 12),
                  if (_nextDayErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _nextDayErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
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

                  ElevatedButton(
                    onPressed: _hello,
                    child: const Text('夢の世界に呼びかける'),
                  ),

                  if (_message != null) const SizedBox(height: 12),
                  if (_message != null)
                    TranslucentPanel(child: Text(_message!)),

                  if (_helloErrorMessage != null) const SizedBox(height: 12),
                  if (_helloErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _helloErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

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
