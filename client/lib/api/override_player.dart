import 'package:client/api/post.dart';

Future<void> apiOverridePlayer(String playerId) async {
  await apiPost('players/$playerId/override', {});
}
