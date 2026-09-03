import 'package:client/api/get.dart';

Future<Map<String, dynamic>> apiListActions(String playerId) async {
  return apiGet('players/$playerId/actions');
}
