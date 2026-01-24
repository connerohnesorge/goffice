# Add Word Comments Enhancements

## Overview
Enhance Word comments support with threaded replies, comment resolution, mention support, and comment range highlighting to match modern Office capabilities.

## Motivation
Current goffice has basic comment support but lacks modern features like replies, resolution state, and mentions needed for collaboration workflows.

## Goals
- Implement comment threading (reply to comments, nested replies, reply metadata)
- Implement comment resolution (mark done/resolved, track resolution author/timestamp, filter by state)
- Support @mentions in comments (parse syntax, link to user directory, notification hooks)
- Enhance comment ranges (multi-paragraph, table cells, image/shape comments)
- Implement comment management (get all, by author/date, delete with replies, move/reanchor)
- Support comment styling and formatting

## Dependencies
- Depends on: wordprocessing core, current comment implementation
- Related: add-change-tracking-implementation

## Estimated Effort
4 weeks

## Priority
P1
