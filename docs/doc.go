// file 'pkg.davsk.net/frodo/doc.go

// by David Lynn Skinner
// on April 2, 2021
// for Davsk Ltd Co

// Copyright 2021 Davsk Ltd Co. All rights reserved.
// Use of this source code is governed by an MIT style license that can be found
// in the LICENSE file.

/*
Package frodo implements Foundational Recursive Orthogonal Development Orders in
a wise way for Davsk Ltd Co CASE tools. First run acquires user info, verifies
dev env and installs CASE tools. Program may be run again to perform specific
actions.

CONCEPT

Projects team members have spent a great deal of time configuring their
respective environments and toolchains. Frodo brings a certain amount of
standardization in procedures and policies through the automation of the
implementation of organization standards. Ideally an individual may implement
different organization standards on their own work station simultaneously or even
remotely via ssh for transient assignments.

OVERVIEW

Frodo may be deployed on any platform as a CLI or GUI app. It is used as a
wrapper for a variety of CASE tools. It is very much a dynamic work in progress
and should not be considered stable at this time.

User Info

User profile:
    * User Name
    * User Email
    * User FullName

Verify Environment

Install CASE Tools

Update Toolchain

New Package 'foo'

Frodo can start a new package for you.
    * mkdir foo
    * cd foo
    * go mod init pkg.davsk.net/foo
    * git init

BUG(david@davsk.net): This is an ALPHA package under development.
*/
package frodo // import pkg.davsk.net/frodo
