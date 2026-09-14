import 'package:flutter/material.dart';

import 'package:client/const.dart';
import 'package:client/screens/create_screen.dart';
import 'package:client/screens/debug_screen.dart';
import 'package:client/screens/explore_screen.dart';
import 'package:client/screens/front_screen.dart';
import 'package:client/screens/home_screen.dart';
import 'package:client/screens/key_screen.dart';
import 'package:client/screens/write_screen.dart';
import 'package:client/state/api_state.dart';
import 'package:client/state/app_state.dart';
import 'package:client/widgets/translucent_panel.dart';

IconData? apiStateIconData() {
  switch (apiState.value) {
    case .disconnected:
      return Icons.cloud_off;
    case .unauthorized:
      return Icons.lock_outline;
    case .serverError:
      return Icons.error_outline;
    default:
      return null;
  }
}

class App extends StatelessWidget {
  final AppState state;

  const App({super.key, required this.state});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: appTitle,

      theme: ThemeData(
        colorScheme: .fromSeed(seedColor: Colors.purple, brightness: .light),
      ),
      darkTheme: ThemeData(
        colorScheme: .fromSeed(seedColor: Colors.purple, brightness: .dark),
      ),
      themeMode: .system,

      routes: {
        '/': (context) => FrontScreen(state: state),
        '/home': (context) => HomeScreen(state: state),
        '/key': (context) => KeyScreen(state: state),

        '/create': (context) => CreateScreen(state: state),
        '/write': (context) => WriteScreen(state: state),
        '/explore': (context) => ExploreScreen(state: state),

        '/debug': (context) => DebugScreen(state: state),
      },
      initialRoute: '/',

      builder: (context, child) {
        return Stack(
          children: [
            child!,

            ValueListenableBuilder(
              valueListenable: apiState,
              builder: (context, state, _) {
                if (state == .connected) {
                  return const SizedBox.shrink();
                }

                return Positioned(
                  top: 16,
                  right: 16,
                  child: TranslucentPanel(
                    child: Icon(apiStateIconData(), size: 24),
                  ),
                );
              },
            ),
          ],
        );
      },
    );
  }
}
