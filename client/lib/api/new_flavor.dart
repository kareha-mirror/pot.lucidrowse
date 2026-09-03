import 'package:client/api/post.dart';
import 'package:client/models/flavor.dart';

Future<Flavor> apiNewFlavor(String input) async {
  final result = await apiPost('players/flavor', {'input': input});
  return Flavor.fromJson(result);
}
