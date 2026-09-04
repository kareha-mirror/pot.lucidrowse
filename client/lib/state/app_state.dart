import 'package:client/api/load_day.dart';
import 'package:client/api/load_player.dart';
import 'package:client/api/load_user.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';
import 'package:client/models/user.dart';

class AppState {
  int? day;
  User user = User();
  Player? player;

  bool debug = false;

  void clear() {
    day = null;
    player = null;
  }

  Future<bool> sync() async {
    final int newDay = await apiLoadDay();
    if (newDay == day) {
      return false;
    }
    day = newDay;

    user = await apiLoadUser();
    player = await apiLoadPlayer();

    return true;
  }

  Flavor? get flavor => player?.flavor;
  PlayerAction? get action => player?.action;

  bool get authorized => user.authorized;
  int get restAiCalls => 10 - user.aiCalls;
  bool get inhabitant => player?.flavor != null;
  bool get committed => player?.action != null;

  void incrementAiCalls() {
    user.aiCalls++;
  }
}
