import 'package:client/api/get.dart';

Future<String?> apiRegionState(String regionCode) async {
  final result = await apiGet('regions/$regionCode/state');
  return result['state'];
}
