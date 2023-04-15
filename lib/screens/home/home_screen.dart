import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:yoga_training_app/constants/constants.dart';
import 'package:yoga_training_app/screens/home/components/courses.dart';
import 'package:yoga_training_app/screens/home/components/custom_app_bar.dart';
import 'package:yoga_training_app/screens/home/components/diff_styles.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  int selsctedIconIndex = 2;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        extendBody: true,
        body: Padding(
          padding: EdgeInsets.only(top: appPadding * 2),
          child: Column(
            children: [
              CustomAppBar(),
              DiffStyles(),
              Courses(),
            ],
          ),
        ),
      ),
    );
  }
}
