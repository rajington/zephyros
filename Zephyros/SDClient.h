//
//  SDClientInterface.h
//  Zephyros
//
//  Created by Steven on 8/11/13.
//  Copyright (c) 2013 Giant Robot Software. All rights reserved.
//

#import <Foundation/Foundation.h>

@protocol SDClientDelegate <NSObject>

- (void) sendResponse:(id)msg;

@end


@interface SDClient : NSObject

- (void) handleRequest:(NSArray*)msg;

- (void) destroy;

@property (weak) id<SDClientDelegate> delegate;

@end
