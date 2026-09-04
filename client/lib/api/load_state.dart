import 'package:client/api/get.dart';

Future<Map<String, dynamic>> apiLoadState() async {
  return apiGet('state');
}
