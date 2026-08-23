import 'package:flutter/material.dart';

import 'package:client/screens/help_screen.dart';
import 'package:client/screens/read_screen.dart';

class FrontScreen extends StatefulWidget {
  const FrontScreen({super.key});

  @override
  State<FrontScreen> createState() => _FrontScreenState();
}

class _FrontScreenState extends State<FrontScreen> {
  bool _ready = false;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _loadBackground();

    precacheImage(const AssetImage('assets/images/home.webp'), context);
    precacheImage(const AssetImage('assets/images/night.webp'), context);
    precacheImage(const AssetImage('assets/images/help.webp'), context);
    precacheImage(const AssetImage('assets/images/debug.webp'), context);
  }

  Future<void> _loadBackground() async {
    await precacheImage(const AssetImage('assets/images/front.webp'), context);
    await precacheImage(const AssetImage('assets/images/logo.webp'), context);

    if (!mounted) return;

    setState(() => _ready = true);
  }

  @override
  Widget build(BuildContext context) {
    if (!_ready) {
      return const Scaffold(body: Center(child: CircularProgressIndicator()));
    }

    return Scaffold(
      body: Stack(
        children: [
          SizedBox(
            width: double.infinity,
            height: double.infinity,
            child: const Image(
              image: AssetImage('assets/images/front.webp'),
              fit: BoxFit.cover,
            ),
          ),

          Center(
            child: Padding(
              padding: EdgeInsets.all(
                MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
              ),
              child: ConstrainedBox(
                constraints: const BoxConstraints(maxWidth: 320),
                child: Column(
                  children: [
                    Image(image: AssetImage('assets/images/logo.webp')),
                  ],
                ),
              ),
            ),
          ),

          Center(
            child: Column(
              mainAxisAlignment: .center,
              children: [
                ElevatedButton(
                  onPressed: () => Navigator.pushNamed(context, '/home'),
                  child: const Text('古びた宿'),
                ),

                const SizedBox(height: 32),

                helpButton(context, HelpPage.front),
              ],
            ),
          ),

          Positioned(
            left: 0,
            right: 0,
            bottom: MediaQuery.paddingOf(context).bottom + 24,
            child: Center(
              child: ElevatedButton(
                onPressed: () => Navigator.pushNamed(context, '/debug'),
                child: const Text('汚れた道具箱'),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
