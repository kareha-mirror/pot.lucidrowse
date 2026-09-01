class Flavor {
  String input = '';

  String name = '';
  String race = '';
  String job = '';

  String description = '';

  String areaCode = '';
  String areaName = '';

  String? imageId;

  Flavor();

  factory Flavor.fromJson(Map<String, dynamic> json) {
    final flavor = Flavor();

    flavor.name = json['name'] as String;
    flavor.race = json['race'] as String;
    flavor.job = json['job'] as String;
    flavor.description = json['description'] as String;
    flavor.areaCode = json['area-code'] as String;
    flavor.areaName = json['area-name'] as String;

    return flavor;
  }

  bool get hasDescription => description != '';

  Flavor.copy(Flavor other)
    : input = other.input,
      name = other.name,
      race = other.race,
      job = other.job,
      description = other.description,
      areaCode = other.areaCode,
      areaName = other.areaName,
      imageId = other.imageId;

  String formatText() {
    final buf = StringBuffer();

    buf.write('名前: $name\n');
    buf.write('種族: $race\n');
    buf.write('職業: $job\n');
    buf.write('\n');
    buf.write('特徴の説明: $description\n');
    buf.write('\n');
    //buf.write('住んでいる地域のコード: $areaCode\n');
    buf.write('住んでいる地域: $areaName\n');

    return buf.toString();
  }
}

class PlayerAction {
  String input = '';

  String description = '';

  String? imageId;

  String date = '';

  bool get hasDescription => description != '';
}

class Player {
  bool inhabitant = false;
  String id = '';

  Flavor flavor = Flavor();

  PlayerAction action = PlayerAction();
  bool committed = false;
}
