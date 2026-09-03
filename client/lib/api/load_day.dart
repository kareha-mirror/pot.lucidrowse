import 'package:client/api/get.dart';

Future<int> apiLoadDay() async {
  final result = await apiGet('day');
  return result['day'];
}
