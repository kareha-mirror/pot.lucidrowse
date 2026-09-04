import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';
import 'package:client/state/api_state.dart';

Future<Map<String, dynamic>> apiPost<T>(
  String path,
  Map<String, dynamic> req,
) async {
  late http.Response response;

  try {
    response = await http.post(
      Uri.parse('$apiBase/$path'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode(req),
    );
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
