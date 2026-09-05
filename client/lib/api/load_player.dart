import 'package:client/api/get.dart';
import 'package:client/models/action.dart';
import 'package:client/models/flavor.dart';
import 'package:client/models/player.dart';

Future<Player?> apiLoadPlayer() async {
  final result = await apiGet('players');

  if (result['id'] == null) {
    return null;
  }

  final player = Player(id: result['id'], day: result['day']);
  player.points = result['points'];
  player.pointsToUpdate = result['points-to-update'];

  if (result['flavor'] != null) {
    final flavor = Flavor.fromJson(result['flavor']);
    flavor.imageId = result['flavor-image-id'];
    player.flavor = flavor;
  }

  if (result['action'] != null) {
    final action = PlayerAction();
    action.description = result['action']['description'];
    action.imageId = result['action-image-id'];
    player.action = action;
  }

  return player;
}
