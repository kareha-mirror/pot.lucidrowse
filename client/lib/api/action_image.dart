import 'dart:convert';
import 'dart:typed_data';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

Future<Uint8List> apiActionImage(String id) async {
  try {
    final uri = Uri.parse('http://localhost:8080/api/player/action-image');

    final response = await http.post(
      uri,
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'id': id}),
    );

    if (response.statusCode != 200) {
      print('HTTP error: ${response.statusCode}');
      return [] as Uint8List;
    }

    return response.bodyBytes as Uint8List;
  } catch (e) {
    print('Connection error: $e');
  }
  return [] as Uint8List;
}
