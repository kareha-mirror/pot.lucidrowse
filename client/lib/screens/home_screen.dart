import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';

class HomeScreen extends StatefulWidget {
  final AppState state;

  const HomeScreen({super.key, required this.state});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('古びた宿'),
      ),

      body: Stack(
        children: [
          Center(
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 1200),
              child: const Image(image: AssetImage('assets/images/home.webp')),
            ),
          ),

          Center(
            child: SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
              child: Column(
                mainAxisAlignment: .center,
                children: [
                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/create'),
                    child: const Text('まどろみの水晶球'),
                  ),

                  const SizedBox(height: 48),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/write'),
                    child: const Text('やわらかい羽ペン'),
                  ),

                  const SizedBox(height: 24),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/read'),
                    child: const Text('おしゃれな日記帳'),
                  ),

                  const SizedBox(height: 24),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/explore'),
                    child: const Text('光おどる小箱'),
                  ),

                  const SizedBox(height: 48),

                  ElevatedButton(
                    onPressed: () => Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) =>
                            const HelpScreen(page: HelpPage.home),
                      ),
                    ),
                    child: const Text('親切な鏡の精'),
                  ),

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: const Text('外を見る'),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
