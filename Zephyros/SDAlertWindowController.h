//
//  SDAlertWindowController.h
//  Zephyros
//
//  Created by Steven on 4/14/13.
//  Copyright (c) 2013 Giant Robot Software. All rights reserved.
//

#import <Cocoa/Cocoa.h>

@interface SDAlertWindowController : NSWindowController

+ (SDAlertWindowController*) sharedAlertWindowController;

- (void) show:(NSString*)oneLineMsg delay:(NSNumber*)delay;

@property CGFloat alertDisappearDelay;
- (void) setAlertAnimates:(BOOL)alertAnimates;

@end
