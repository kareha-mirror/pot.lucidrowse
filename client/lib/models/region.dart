class Area {
  final String name;

  const Area({required this.name});
}

class Region {
  final String code;
  final String name;
  final List<Area> areas;
  final String image;

  const Region({
    required this.code,
    required this.name,
    required this.areas,
    required this.image,
  });
}

const List<Region> regions = [
  Region(
    code: 'capital',
    name: '王都地方',
    areas: [
      Area(name: '王都'),
      Area(name: '近郊'),
    ],
    image: 'assets/images/regions/capital.webp',
  ),
  Region(
    code: 'west',
    name: '西方地方',
    areas: [
      Area(name: '街道'),
      Area(name: '丘陵'),
    ],
    image: 'assets/images/regions/west.webp',
  ),
  Region(
    code: 'north',
    name: '北方地方',
    areas: [
      Area(name: '山岳'),
      Area(name: '高原'),
    ],
    image: 'assets/images/regions/north.webp',
  ),
  Region(
    code: 'east',
    name: '東方地方',
    areas: [
      Area(name: '森林'),
      Area(name: '湖沼'),
    ],
    image: 'assets/images/regions/east.webp',
  ),
  Region(
    code: 'south',
    name: '南方地方',
    areas: [
      Area(name: '平野'),
      Area(name: '農村'),
    ],
    image: 'assets/images/regions/south.webp',
  ),
  Region(
    code: 'coast',
    name: '海岸地方',
    areas: [
      Area(name: '港湾'),
      Area(name: '海岸'),
    ],
    image: 'assets/images/regions/coast.webp',
  ),
  Region(
    code: 'islands',
    name: '島嶼地方',
    areas: [
      Area(name: '大島'),
      Area(name: '周辺諸島'),
    ],
    image: 'assets/images/regions/islands.webp',
  ),
];
