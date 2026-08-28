import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:client/const.dart';

Future<Map<String, dynamic>> apiRegionState(String regionId) async {
  final response = await http.get(Uri.parse('$apiBase/region-state/$regionId'));

  if (response.statusCode != 200) {
    throw Exception('HTTP error: ${response.statusCode}');
  }

  return jsonDecode(response.body);
}
