import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';

Future<Map<String, dynamic>> apiUpdateFlavor(String input) async {
  final response = await http.post(
    Uri.parse('$apiBase/players/flavor/update'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({'input': input}),
  );

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
