import 'package:flutter/material.dart';

import 'package:client/state/app_state.dart';
import 'package:client/utils/game_calendar.dart';

class ReadScreen extends StatefulWidget {
  final AppState state;

  const ReadScreen({super.key, required this.state});

  @override
  State<ReadScreen> createState() => _ReadScreenState();
}

class _ReadScreenState extends State<ReadScreen> {
  List<Widget> _actionList() {
    if (widget.state.player.actions.isEmpty) {
      return [const SizedBox(height: 24), const Text('まだ何も書かれてない。')];
    }

    List<Widget> list = [];
    for (final action in widget.state.player.actions) {
      list.add(const SizedBox(height: 24));
      list.add(Text(formatGameDate(action.day)));
      list.add(Text(action.filtered));
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('おしゃれな日記帳'),
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
        child: Center(
          child: Column(
            children: [
              ..._actionList(),
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('閉じる'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
