import 'package:client/api/post.dart';
import 'package:client/models/flavor.dart';

Future<Flavor> apiUpdateFlavor(String input) async {
  final result = await apiPost('players/flavor/update', {'input': input});
  return Flavor.fromJson(result);
}
