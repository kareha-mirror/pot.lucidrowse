class User {
  bool authorized = false;
  String? name;
  int aiCalls = 0;

  User();

  factory User.fromJson(Map<String, dynamic> json) {
    final user = User();

    user.authorized = json['authorized'] as bool;
    user.name = json['name'] as String?;
    user.aiCalls = json['ai-calls'] as int;

    return user;
  }
}
