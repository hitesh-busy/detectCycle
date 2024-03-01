## Problem Statement

Detect if there is a cycle in a directed graph representing employee-mentor relationships. If no cycle is detected, update the database table.

## Overview

This project aims to solve the problem of detecting cycles in a table that represents the relationships between employees and their mentors. The relationships are modeled as a directed graph, and we use Depth-First Search (DFS) to identify cycles.

## Implementation

We utilize a DFS-based approach to traverse the graph and detect cycles. Two slices are used to keep track of visited nodes and the current DFS path. If a cycle is detected, appropriate actions are taken.

## Example Tree
 1     2
   \   /
    3 -- 6
   / \
  4   5

The arrows are pointing downwards to represent the direction. now  3 -> 6 and 3->5 seem like a cycle but its not.
