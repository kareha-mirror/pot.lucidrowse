import 'package:client/api/post.dart';
import 'package:client/models/action.dart';

Future<PlayerAction> apiNewAction(String input) async {
  final result = await apiPost('players/actions', {'input': input});
  return PlayerAction.fromJson(result);
}
