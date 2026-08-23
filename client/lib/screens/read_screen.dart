import 'package:flutter/material.dart';

import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/widgets/translucent_panel.dart';

class ReadScreen extends StatefulWidget {
  final AppState state;

  final Player player;

  const ReadScreen({super.key, required this.state, required this.player});

  @override
  State<ReadScreen> createState() => _ReadScreenState();
}

class _ReadScreenState extends State<ReadScreen> {
  List<Widget> _actionList() {
    if (widget.player.actions.isEmpty) {
      return [
        const SizedBox(height: 24),
        TranslucentPanel(child: const Text('まだ何も書かれてない。')),
      ];
    }

    List<Widget> list = [];
    for (final action in widget.player.actions) {
      list.add(const SizedBox(height: 24));
      list.add(
        TranslucentPanel(
          child: Column(
            children: [
              Text(formatDate(action.day)),
              SizedBox(height: 8),
              Text(action.filtered),
            ],
          ),
        ),
      );
      list.add(SizedBox(height: 8));
      if (action.image != null) {
        list.add(
          ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 300),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Image.memory(action.image!),
            ),
          ),
        );
      }
    }
    return list;
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

                  SizedBox(height: 96),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

Widget myDiaryButton(BuildContext context, AppState state) {
  return ElevatedButton(
    onPressed: () => Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ReadScreen(state: state, player: state.player),
      ),
    ),
    child: const Text('おしゃれな日記帳'),
  );
}

Widget diaryButton(BuildContext context, AppState state, Player player) {
  return ElevatedButton(
    onPressed: () => Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ReadScreen(state: state, player: player),
      ),
    ),
    child: const Text('日記'),
  );
}
