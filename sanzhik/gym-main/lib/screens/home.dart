import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:yoga_training_app/constants/constants.dart';
import 'package:yoga_training_app/screens/home/components/courses.dart';
import 'package:yoga_training_app/screens/home/components/custom_app_bar.dart';
import 'package:yoga_training_app/screens/home/components/diff_styles.dart';
import 'package:yoga_training_app/screens/home/home_screen.dart';
import 'package:yoga_training_app/screens/profile/profile.dart';
import 'package:yoga_training_app/screens/search/search.dart';

import 'favs/favs.dart';
import 'media/media.dart';

class Home extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<Home> {
  List pages = [Media(), Search(), HomeScreen(), Favs(), Profile()];
  int currentPage = 2;
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        extendBody: true,
        body: pages[currentPage],
        bottomNavigationBar: CurvedNavigationBar(
          backgroundColor: Colors.transparent,
          index: currentPage,
          buttonBackgroundColor: primary,
          height: 60.0,
          color: white,
          onTap: (index) {
            setState(() {
              currentPage = index;
            });
          },
          animationDuration: Duration(
            milliseconds: 200,
          ),
          items: <Widget>[
            Icon(
              Icons.play_arrow_outlined,
              size: 30,
              color: currentPage == 0 ? white : black,
            ),
            Icon(
              Icons.search,
              size: 30,
              color: currentPage == 1 ? white : black,
            ),
            Icon(
              Icons.home_outlined,
              size: 30,
              color: currentPage == 2 ? white : black,
            ),
            Icon(
              Icons.favorite_border_outlined,
              size: 30,
              color: currentPage == 3 ? white : black,
            ),
            Icon(
              Icons.person_outline,
              size: 30,
              color: currentPage == 4 ? white : black,
            ),
          ],
        ),
      ),
    );
  }
}
