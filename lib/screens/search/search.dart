import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:yoga_training_app/constants/constants.dart';
import 'package:yoga_training_app/screens/home/components/courses.dart';
import 'package:yoga_training_app/screens/home/components/custom_app_bar.dart';
import 'package:yoga_training_app/screens/home/components/diff_styles.dart';

class Search extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<Search> {
  int selsctedIconIndex = 2;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        extendBody: true,
        body: Padding(
          padding: const EdgeInsets.only(top: appPadding * 2),
          child: Column(
            children: [
              Container(
                padding: EdgeInsets.all(8),
                margin: EdgeInsets.all(16),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(16),
                  color: Color.fromARGB(255, 236, 236, 234),
                ),
                child: Row(
                  children: const [
                    Expanded(
                      child: TextField(
                        decoration: InputDecoration(
                            focusColor: Colors.black,
                            border: InputBorder.none,
                            hintText: "Search",
                            hintStyle:
                                TextStyle(color: Colors.grey, fontSize: 15)),
                      ),
                    ),
                    Icon(Icons.search, color: Colors.grey),
                  ],
                ),
              ),
              // CustomAppBar(),
              DiffStyles(),
              Courses(),
            ],
          ),
        ),
      ),
    );
  }
}
