import 'package:client/api/post.dart';

Future<String> apiImageFlavor() async {
  final result = await apiPost('players/flavor/image', {});
  return result['image-id'];
}
