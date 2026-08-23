import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';
import 'package:client/screens/read_screen.dart';
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
            padding: EdgeInsets.all(
              MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
            ),
            child: Center(
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

                  const SizedBox(height: 32),

                  myDiaryButton(context, widget.state),

                  const SizedBox(height: 32),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/explore'),
                    child: const Text('光おどる小箱'),
                  ),

                  const SizedBox(height: 48),

                  helpButton(context, HelpPage.home),

                  const SizedBox(height: 64),

                  ElevatedButton(
                    onPressed: () => Navigator.pushNamed(context, '/'),
                    child: const Text('外を見る'),
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
