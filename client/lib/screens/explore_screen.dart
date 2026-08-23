import 'package:flutter/material.dart';

import 'package:client/models/region.dart';
import 'package:client/state/app_state.dart';
import 'package:client/widgets/region_card.dart';
import 'package:client/widgets/translucent_panel.dart';

class ExploreScreen extends StatefulWidget {
  final AppState state;

  const ExploreScreen({super.key, required this.state});

  @override
  State<ExploreScreen> createState() => _ExploreScreenState();
}

class _ExploreScreenState extends State<ExploreScreen> {
  int regionIndex = -1;

  List<Widget> inhabitantsList() {
    List<Widget> list = [];
    for (final inhabitant in widget.state.inhabitants) {
      final flavor = inhabitant.lastFlavor;

      list.add(const SizedBox(height: 24));

      if (flavor.image != null) {
        list.add(
          ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 300),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Image.memory(flavor.image!),
            ),
          ),
        );
      }

      list.add(const SizedBox(height: 8));

      list.add(
        TranslucentPanel(
          child: Text('${flavor.name} / ${flavor.race} / ${flavor.job}'),
        ),
      );

      list.add(const SizedBox(height: 8));

      list.add(TranslucentPanel(child: Text(flavor.filtered)));
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          if (regionIndex == -1)
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(
                  widget.state.player.committed
                      ? 'assets/images/night.webp'
                      : 'assets/images/home.webp',
                ),
                fit: BoxFit.cover,
              ),
            )
          else
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(regions[regionIndex].image),
                fit: BoxFit.cover,
              ),
            ),

          if (regionIndex == -1)
            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
              child: Center(
                child: Column(
                  children: [
                    const SizedBox(height: 48),

                    TranslucentPanel(child: const Text('夢の世界の地域たち')),

                    const SizedBox(height: 12),

                    for (int index = 0; index < regions.length; index++)
                      Padding(
                        padding: EdgeInsets.symmetric(
                          horizontal: MediaQuery.sizeOf(context).width < 600
                              ? 6
                              : 48,
                          vertical: MediaQuery.sizeOf(context).width < 600
                              ? 3
                              : 24,
                        ),
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 600),
                          child: RegionCard(
                            region: regions[index],
                            onTap: () => setState(() => regionIndex = index),
                          ),
                        ),
                      ),

                    ElevatedButton(
                      onPressed: () => Navigator.pop(context),
                      child: const Text('ふたを閉じる'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            )
          else
            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
              child: Center(
                child: Column(
                  children: [
                    const SizedBox(height: 48),

                    TranslucentPanel(child: const Text('夢の世界の住人たち')),

                    const SizedBox(height: 24),

                    if (widget.state.inhabitants.isEmpty)
                      TranslucentPanel(child: const Text('まだ誰も住んでない。'))
                    else
                      ...inhabitantsList(),

                    const SizedBox(height: 96),

                    ElevatedButton(
                      onPressed: () => setState(() => regionIndex = -1),
                      child: const Text('中を見直す'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            ),
        ],
      ),
    );
  }
}
