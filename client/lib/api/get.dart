import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';
import 'package:client/state/api_state.dart';

Future<Map<String, dynamic>> apiGet<T>(String path) async {
  late http.Response response;

  try {
    response = await http.get(Uri.parse('$apiBase/$path'));
  } catch (e) {
    apiState.value = ApiState.disconnected;
    rethrow;
  }

  if (response.statusCode == 502) {
    apiState.value = ApiState.disconnected;
    throw Exception('HTTP error: ${response.statusCode}');
  }

  if (response.statusCode == 401) {
    apiState.value = ApiState.unauthorized;
    throw Exception('HTTP error: ${response.statusCode}');
  }

  if (response.statusCode != 200) {
    apiState.value = ApiState.serverError;
    throw Exception('HTTP error: ${response.statusCode}');
  }

  apiState.value = ApiState.connected;
  return jsonDecode(response.body);
}
