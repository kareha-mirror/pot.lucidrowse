import 'package:client/api/get.dart';

Future<Map<String, dynamic>> apiListPlayers(String regionCode) async {
  return apiGet('regions/$regionCode/players');
}
