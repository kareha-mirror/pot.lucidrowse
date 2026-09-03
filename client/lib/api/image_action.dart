import 'package:client/api/post.dart';

Future<String> apiImageAction() async {
  final result = await apiPost('players/actions/image', {});
  return result['image-id'];
}
