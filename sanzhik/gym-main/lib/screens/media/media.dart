import 'dart:convert';

import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:yoga_training_app/constants/constants.dart';
import 'package:yoga_training_app/screens/home/components/courses.dart';
import 'package:yoga_training_app/screens/home/components/custom_app_bar.dart';
import 'package:yoga_training_app/screens/home/components/diff_styles.dart';
import 'package:http/http.dart' as http;
import '../../widgets/MusicCard.dart';

class Media extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<Media> {
  int selsctedIconIndex = 2;
  List<dynamic> artists = [];
  // fetchArtists();

  void initState() {
    fetchArtists();
  }

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        extendBody: true,
        body: Padding(
          padding: const EdgeInsets.symmetric(
            horizontal: appPadding * 1,
            vertical: appPadding * 1,
          ),
          child:
              Column(crossAxisAlignment: CrossAxisAlignment.start, children: [
            Container(
              margin: EdgeInsets.only(bottom: 24),
              child: const Text(
                'Overlay display',
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.w800,
                  letterSpacing: 1.5,
                ),
              ),
            ),
            Expanded(
              child: ListView(
                children: [
                  if (artists.length != 0)
                    for (var artist in artists)
                      Container(
                        child: MusicCard(
                          name: artist["name"],
                          image: artist["image"],
                        ),
                      ),
                  if (artists.length == 0)
                    Container(
                      margin: EdgeInsets.only(top: 48),
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                        children: [
                          CircularProgressIndicator(), //<-- SEE HERE
                        ],
                      ),
                    )
                ],
              ),
            ),
          ]),
        ),
      ),
    );
  }

  Future<void> fetchArtists() async {
    try {
      final response = await http
          .get(Uri.parse('https://groupietrackers.herokuapp.com/api/artists'));

      if (response.statusCode == 200) {
        var tagsJson = jsonDecode(response.body);
        List<dynamic> tags = List.from(tagsJson);
        setState(() {
          artists = tags;
        });
      } else {
        throw Exception('Failed to load album');
      }
    } catch (e) {}
  }
}

// class Data {
//   Artist artists;

//   Data({
//     required this.artists,
//   });

//   factory Data.fromJson(Map<String, dynamic> json) {
//     return Data(
//       artists: Artist.fromJson(json["artists"]),
//     );
//   }
// }

// class Artist {
//   int id;
//   String name;
//   String image;

//   Artist({
//     required this.id,
//     required this.name,
//     required this.image,
//   });

//   factory Artist.fromJson(Map<String, dynamic> json) {
//     return Artist(
//       id: json["id"],
//       name: json["name"],
//       image: json["image"],
//     );
//   }
// }

