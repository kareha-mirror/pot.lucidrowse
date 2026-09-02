import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';

Future<Map<String, dynamic>> apiLoadFlavor() async {
  final response = await http.get(Uri.parse('$apiBase/players/flavor'));

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
