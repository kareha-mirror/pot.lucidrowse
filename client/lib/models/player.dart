enum InhabitKind {
  foreigner,
  inhabitant,
  hermit,
  forgotten;

  bool get isForeigner => this == InhabitKind.foreigner;
}

class Player {
  InhabitKind inhabitKind;
  String flavorText;
  String actionText;

  Player.unnamed()
    : inhabitKind = InhabitKind.foreigner,
      flavorText = '',
      actionText = '';
}
