import 'package:flutter/material.dart';

import 'package:client/state/app_state.dart';

class ExploreScreen extends StatefulWidget {
  final AppState state;

  const ExploreScreen({super.key, required this.state});

  @override
  State<ExploreScreen> createState() => _ExploreScreenState();
}

class _ExploreScreenState extends State<ExploreScreen> {
  List<Widget> inhabitantsList() {
    List<Widget> list = [];
    for (final inhabitant in widget.state.inhabitants) {
      list.add(SizedBox(height: 24));
      list.add(
        ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 300),
          child: Image(image: AssetImage(inhabitant.flavor.imageUrl)),
        ),
      );
      list.add(
        Text(
          '${inhabitant.flavor.name} / ${inhabitant.flavor.race} / ${inhabitant.flavor.job}',
        ),
      );
      list.add(Text(inhabitant.action.filtered));
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('光おどる小箱'),
      ),
      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
          child: Column(
            mainAxisAlignment: .center,
            children: [
              SizedBox(height: 48),
              Text('夢の世界の住人たち'),
              if (widget.state.inhabitants.isEmpty) ...[
                SizedBox(height: 48),
                Text('まだ誰も住んでない。'),
              ] else ...[
                SizedBox(height: 24),
                ...inhabitantsList(),
              ],
              SizedBox(height: 96),
              ElevatedButton(
                onPressed: () => Navigator.pop(context),
                child: Text('ふたを閉じる'),
              ),
              SizedBox(height: 48),
            ],
          ),
        ),
      ),
    );
  }
}
