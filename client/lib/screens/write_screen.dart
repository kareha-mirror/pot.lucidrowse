import 'package:flutter/material.dart';

import 'package:client/api/commit_action.dart';
import 'package:client/api/day.dart';
import 'package:client/api/ensure_session.dart';
import 'package:client/api/load_action.dart';
import 'package:client/api/load_flavor.dart';
import 'package:client/api/new_action.dart';
import 'package:client/api/image_action.dart';
import 'package:client/models/player.dart';
import 'package:client/state.dart';
import 'package:client/utils/calendar.dart';
import 'package:client/utils/image_url.dart';
import 'package:client/widgets/translucent_panel.dart';

class WriteScreen extends StatefulWidget {
  final AppState state;

  const WriteScreen({super.key, required this.state});

  @override
  State<WriteScreen> createState() => _WriteScreenState();
}

class _WriteScreenState extends State<WriteScreen> {
  Flavor _committedFlavor = Flavor();
  PlayerAction _committedAction = PlayerAction();
  PlayerAction _editingAction = PlayerAction();

  int _day = 0;
  bool _dayLoaded = false;
  String? _dayErrorMessage;
  late TextEditingController _inputController;
  String? _actionErrorMessage;
  late TextEditingController _outputController;
  bool _actionLoading = false;
  bool _imageLoading = false;
  String? _imageErrorMessage;

  @override
  void initState() {
    super.initState();
    _inputController = TextEditingController();
    _outputController = TextEditingController();
  }

  @override
  void dispose() {
    _inputController.dispose();
    _outputController.dispose();
    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _loadCommittedFlavor();
    _loadCommittedAction();
    _loadDay();
  }

  Future<void> _loadCommittedFlavor() async {
    try {
      final result = await apiLoadFlavor();

      final flavor = Flavor.fromJson(result['flavor']);
      flavor.imageId = result['image-id'];

      setState(() {
        _committedFlavor = flavor;
      });
    } catch (e) {
      // ignore
    }
  }

  Future<void> _loadCommittedAction() async {
    try {
      final result = await apiLoadAction();

      final action = PlayerAction();
      action.description = result['action']['description'];
      action.imageId = result['image-id'];

      setState(() {
        _committedAction = action;
        _editingAction = action;
        _outputController.text = action.description;
      });
    } catch (e) {
      // ignore
    }
  }

  Future<void> _loadDay() async {
    setState(() => _dayLoaded = false);

    try {
      final result = await apiDay();

      if (!mounted) return;

      setState(() {
        _day = result['day'];

        _dayLoaded = true;
      });
    } catch (e) {
      setState(() => _dayErrorMessage = e.toString());
    }
  }

  Future<void> _loadImage() async {
    setState(() {
      _imageLoading = true;
      _imageErrorMessage = null;
    });
    try {
      final result = await apiImageAction();

      setState(() {
        _editingAction.imageId = result['image-id'];
      });
    } catch (e) {
      setState(() => _imageErrorMessage = e.toString());
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _loadAction() async {
    setState(() {
      _actionLoading = true;
      _actionErrorMessage = null;
    });
    try {
      await apiEnsureSession();

      final result = await apiNewAction(_inputController.text);
      if (result['error'] == '') {
        final action = PlayerAction();

        action.input = _inputController.text;

        action.description = result['description'];

        setState(() {
          _outputController.text = action.description;
          _editingAction = action;
        });

        _loadImage();
      } else {
        setState(() {
          _outputController.text = 'エラー:\n${result['error']}';
        });
      }
    } catch (e) {
      setState(() => _actionErrorMessage = e.toString());
    } finally {
      setState(() => _actionLoading = false);
    }
  }

  Future<void> _commit() async {
    await apiCommitAction();

    setState(() {
      _committedAction = _editingAction;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          SizedBox(
            width: double.infinity,
            height: double.infinity,
            child: Image(
              /* TODO night
              image: AssetImage(
                _player.committed
                    ? 'assets/images/night.webp'
                    : 'assets/images/home.webp',
              ),
              */
              image: AssetImage('assets/images/home.webp'),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  if (_committedFlavor.hasDescription) ...[
                    if (_dayLoaded)
                      TranslucentPanel(child: Text(formatDate(_day))),

                    if (_dayErrorMessage != null) const SizedBox(height: 12),
                    if (_dayErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _dayErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    const SizedBox(height: 24),

                    TranslucentPanel(child: const Text('日記に書いてみよう。')),

                    const SizedBox(height: 24),

                    TranslucentPanel(child: const Text('今日は何をして過ごしましたか？')),
                  ],

                  const SizedBox(height: 24),

                  if (_committedFlavor.hasDescription)
                    Padding(
                      padding: EdgeInsets.all(
                        MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
                      ),
                      child: TranslucentPanel(
                        child: TextField(
                          controller: _inputController,
                          decoration: const InputDecoration(
                            hintText: '(今日、何をしたか書いてみよう。)',
                          ),
                          onChanged: (String value) => setState(() {}),
                          maxLines: null,
                          maxLength: 140,
                          readOnly:
                              _actionLoading ||
                              _imageLoading ||
                              _committedAction.hasDescription,
                        ),
                      ),
                    ),

                  const SizedBox(height: 24),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed:
                          (_inputController.text.isEmpty ||
                              _actionLoading ||
                              _imageLoading ||
                              _committedAction.hasDescription)
                          ? null
                          : _loadAction,
                      child: const Text('日記に書く'),
                    ),
                  ),

                  if (_actionErrorMessage != null) const SizedBox(height: 12),
                  if (_actionErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _actionErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

                  if (_actionLoading)
                    const CircularProgressIndicator()
                  else if (_editingAction.hasDescription) ...[
                    const SizedBox(height: 48),

                    TranslucentPanel(
                      child: const Text('あなたの言葉は夢に映され、こうなりました。'),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: TextField(
                        controller: _outputController,
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                    const SizedBox(height: 24),

                    if (_imageLoading)
                      const Column(
                        children: [
                          TranslucentPanel(child: Text('情景を映し出しています…')),
                          CircularProgressIndicator(),
                        ],
                      )
                    else if (_editingAction.imageId != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.network(
                              imageUrl(_editingAction.imageId!),
                            ),
                          ),
                        ),
                      ),

                    if (_imageErrorMessage != null) const SizedBox(height: 12),
                    if (_imageErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _imageErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed:
                          (_actionLoading ||
                              _imageLoading ||
                              _committedAction.hasDescription)
                          ? null
                          : () async {
                              await _commit();

                              if (!context.mounted) return;

                              Navigator.pushNamed(context, '/home');
                            },
                      child: const Text('これでよし'),
                    ),
                  ],

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: (_actionLoading || _imageLoading)
                        ? null
                        : () => Navigator.pop(context),
                    child: const Text('ペンを置く'),
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
