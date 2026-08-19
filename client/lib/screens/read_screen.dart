import 'package:flutter/material.dart';

import 'package:client/state/app_state.dart';

class ReadScreen extends StatefulWidget {
  final AppState state;

  const ReadScreen({super.key, required this.state});

  @override
  State<ReadScreen> createState() => _ReadScreenState();
}

class _ReadScreenState extends State<ReadScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('おしゃれな日記帳'),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              SizedBox(height: 48),
              Text('日記帳はまだ開かない。'),
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('閉じたまま'),
              ),
              SizedBox(height: 48),
            ],
          ),
        ),
      ),
    );
  }
}
