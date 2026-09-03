import 'package:flutter/foundation.dart';

enum ApiState { connected, disconnected, unauthorized, serverError }

final apiState = ValueNotifier<ApiState>(ApiState.connected);
