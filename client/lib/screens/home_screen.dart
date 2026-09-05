import 'dart:async';
import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';
import 'package:client/screens/read_screen.dart';
import 'package:client/state/app_state.dart';
import 'package:client/widgets/translucent_panel.dart';

class HomeScreen extends StatefulWidget {
  final AppState state;

  const HomeScreen({super.key, required this.state});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  Timer? _syncTimer;

  @override
  void initState() {
    super.initState();

    _syncTimer = Timer.periodic(const Duration(hours: 1), (_) => _sync());
  }

  @override
  void dispose() {
    _syncTimer?.cancel();

    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _sync();
  }

  Future<void> _sync() async {
    try {
      if (await widget.state.sync()) {
        if (!mounted) return;

        setState(() {});
      }
    } catch (e) {
      // ignore
    }
  }

  @override
  Widget build(BuildContext context) {
    return PopScope(
      canPop: false,
      onPopInvokedWithResult: (didPop, result) {
        if (!didPop) {
          Navigator.pushReplacementNamed(context, '/');
        }
      },
      child: Scaffold(
        body: Stack(
          children: [
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(
                  widget.state.committed
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
                    if (!widget.state.committed && widget.state.updatable)
                      const SizedBox(height: 24),
                    if (!widget.state.committed && widget.state.updatable)
                      TranslucentPanel(child: const Text('何か変化はありましたか？')),
                    if (!widget.state.committed && widget.state.updatable)
                      const SizedBox(height: 24),

                    ElevatedButton(
                      onPressed: () => Navigator.pushNamed(context, '/create'),
                      child: const Text('まどろみの水晶球'),
                    ),

                    if (widget.state.committed) const SizedBox(height: 48),
                    if (widget.state.committed)
                      TranslucentPanel(child: const Text('また明日。')),

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
                      onPressed: () =>
                          Navigator.pushReplacementNamed(context, '/'),
                      child: const Text('外を見る'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            ),

            Positioned(
              bottom: 24,
              left: 8,
              right: 8,
              child: Row(
                children: [
                  for (int i = 0; i < widget.state.restAiCalls; i++)
                    Image(image: AssetImage('assets/images/fruit.webp')),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
