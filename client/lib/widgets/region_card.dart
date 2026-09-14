import 'package:flutter/material.dart';

import 'package:client/models/region.dart';
import 'package:client/widgets/translucent_panel.dart';

class RegionCard extends StatelessWidget {
  const RegionCard({
    super.key,
    required this.region,
    this.playerCount,
    this.topic,
    this.onTap,
  });

  final Region region;
  final int? playerCount;
  final String? topic;
  final GestureTapCallback? onTap;

  @override
  Widget build(BuildContext context) {
    final List<String> areas = [];
    for (final area in region.areas) {
      areas.add(area.name);
    }

    return Column(
      children: [
        Card(
          clipBehavior: .antiAlias,
          child: InkWell(
            onTap: onTap,
            child: Row(
              children: [
                Image.asset(
                  region.thumbnailPath,
                  width: MediaQuery.sizeOf(context).width < 600 ? 150 : 300,
                  height: MediaQuery.sizeOf(context).width < 600 ? 75 : 150,
                  fit: .cover,
                ),

                const SizedBox(width: 16),

                Expanded(
                  child: Column(
                    crossAxisAlignment: .start,
                    children: [
                      Text(
                        region.name,
                        style: Theme.of(context).textTheme.titleMedium,
                      ),

                      const SizedBox(height: 4),

                      Text(areas.join(' / ')),

                      if (playerCount != null) const SizedBox(height: 4),

                      if (playerCount != null) Text('住人 $playerCount'),
                    ],
                  ),
                ),

                const SizedBox(width: 12),
              ],
            ),
          ),
        ),
        if (topic != null) TranslucentPanel(child: Text(topic!)),
      ],
    );
  }
}
