class PlayerAction {
  String input = '';

  String description = '';

  int day = 0;
  String? imageId;

  String error = '';

  PlayerAction();

  factory PlayerAction.fromJson(Map<String, dynamic> json) {
    final action = PlayerAction();

    action.description = json['description'] as String;
    action.error = json['error'] as String;

    return action;
  }

  bool get hasDescription => description != '';
}
