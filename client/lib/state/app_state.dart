import 'package:client/api/api.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';
import 'package:client/models/user.dart';

class AppState {
  String? mode;
  int maxAiCalls = 0;
  int pointsToUpdate = 65535;
  int? day;

  User user = User();
  Player? player;

  bool debug = false;

  void clear() {
    mode = null;
    maxAiCalls = 0;
    pointsToUpdate = 65535;
    day = null;
    user = User();
    player = null;
  }

  Future<bool> sync() async {
    final result = await api.loadState();

    mode = result['mode'];
    maxAiCalls = result['max-ai-calls'];
    pointsToUpdate = result['points-to-update'];

    final newDay = result['day'];
    if (newDay == day) {
      return false;
    }
    day = newDay;

    try {
      user = await api.loadUser();
    } catch (e) {
      // ignore
    }

    try {
      player = await api.loadPlayer();
    } catch (e) {
      // ignore
    }

    return true;
  }

  bool get devel => mode == 'devel';

  Flavor? get flavor => player?.flavor;
  PlayerAction? get action => player?.action;

  bool get authorized => user.authorized;
  int get restAiCalls => maxAiCalls - user.aiCalls;
  bool get inhabitant => player?.flavor != null;
  bool get committed => player?.action != null;

  bool get night => committed || restAiCalls < 1;

  void incrementAiCalls() {
    user.aiCalls++;
  }

  bool get updatable => (player?.points ?? 0) >= pointsToUpdate;
}
