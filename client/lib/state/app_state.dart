import 'package:client/api/api.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';
import 'package:client/models/user.dart';

class AppState {
  String? mode;
  int? day;
  User user = User();
  Player? player;

  bool debug = false;

  void clear() {
    day = null;
    player = null;
  }

  Future<bool> sync() async {
    final result = await api.loadState();
    final newDay = result['day'];
    if (newDay == day) {
      return false;
    }
    day = newDay;

    mode = result['mode'];

    user = await api.loadUser();
    player = await api.loadPlayer();

    return true;
  }

  bool get devel => mode == 'devel';

  Flavor? get flavor => player?.flavor;
  PlayerAction? get action => player?.action;

  bool get authorized => user.authorized;
  int get restAiCalls => user.maxAiCalls - user.aiCalls;
  bool get inhabitant => player?.flavor != null;
  bool get committed => player?.action != null;

  void incrementAiCalls() {
    user.aiCalls++;
  }

  bool get updatable =>
      (player?.points ?? 0) >= (player?.pointsToUpdate ?? 65536);
}
